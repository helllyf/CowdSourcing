pragma solidity >=0.4.21 <0.7.0;

import "./Sortition.sol";
import "./util/Owned.sol";


contract QVote is Sortition ,Owned{

    // ==========
    // EVENTS:
    // ==========

    event VoteCreated(
        address indexed _creator,
        uint _startTime,
        uint _closeTime,
        bytes32 _description
    );


    event VoterApproved(
        address indexed _voter,
        bytes32 _description,
        uint _startTime,
        uint _closeTime
    );

    event VoteCompleted(uint topChoice, address candidate);


    // =========
    // Data Structures and STATE VARIABLES:
    // =========

    enum QVMethod {QV, Scheme1, Scheme2 , Scheme3}

    /// @dev Voting methods are extensible.
    QVMethod qvMethod;
 
    enum VoteStatus {created, commitPhase, revealPhase, completed}
    
    //only two Candidate, worker and requester.
    struct Candidate {
        uint32 voteCount;       // Number of votes revealed for the candidate
        bytes32 name;           // Name of candidate or description of option
        address account;         // Ethereum address account of the candidate
    }
    
    // Circuit-breaker bool
    bool private stopped = false;

    uint createBlock;       // Vote Create Block Number 
    uint32 startTime;       // Earliest time can move into commitPhase
    uint32 closeTime;       // Earliest time can move into revealPhase
    uint32 public firstVoteCost;   // Cost of the first vote (in Wei). Cost for each additional vote scales quadratically
    bytes32 description;    // Brief (32 bytes) description of the Vote
    uint8 candidateCount;  // Number of candidates in the Vote
    VoteStatus status;      // Current status of the Vote. 'created', 'commitPhase', 'revealPhase', or 'completed'
    mapping (uint => Candidate) candidates; // Mapping of candidates in the election
    mapping (address => bool) approvedVoters;   /* Addresses approved
            to vote (to prevent Sybil attacks using multiple addresses
            for cheaper votes) */

    struct Vote {
        uint32 numVotes;        // Number of votes purchased / committed
        /* Commitment hash consisting of number of votes,
        vote (vote is a uint that represents the key in the candidates
        mapping corresponding to the candidate voted for), and salt.
        Note that salt should be kept secret, and a new salt should be
        used each time. */
        bytes32 commitment;
    }

    uint public voterCount = 0;

    // Double mapping from user address => user's vote
    mapping (address => Vote) public votes;
    
    // mapping from user to his candidate
    mapping (address => uint) public voterToCandidate;
    
    /// @dev This is the rate of Sortition
    uint public sortitionRate; 

    // the sortitionRate grow with the blocknumber, 
    uint public rateGrowStep;
    
    // how many people will be selected? 
    uint public voterSelectedNum;

    // ==========
    // MODIFIERS:
    // ==========

    // Circuit-breaker modifiers in case of a problem with the contract
    modifier stopInEmergency { require(!stopped); _; }
    modifier onlyInEmergency { require(stopped); _; }

    modifier commitPhase() {
        /* If it is past vote startTime, but vote is still in Created
        Phase, change to Commit Phase */
        require(
            status == VoteStatus.commitPhase ||
            status == VoteStatus.revealPhase
        );
        _;
    }

    modifier revealPhase() {
        require (
                status == VoteStatus.revealPhase,
                "Vote must be in Reveal phase."
            );
            _;

    }

    modifier completePhase() {
        require(
            status == VoteStatus.completed,
            "vote must be in Completed Phase."
        );
        _;
    }

    modifier onlyApprovedVoters() {
        require (
            approvedVoters[msg.sender] == true,
            "msg.sender must be approved to vote"
        );
        _;
    }

    // ==========
    // VOTEING INTERFACE:
    // ==========

    constructor(address _requester , address _worker )
        public
    {
        require(_requester != address(0));
        require(_worker != address(0));
        addCandidate(bytes32("requester"),_requester);
        addCandidate(bytes32("worker"),_worker);
        
        startTime = uint32(now);
        createBlock = block.number;
        
        // the initial sortition rate of voter is % percent 
        sortitionRate = 10;
        // After create poll game, the sortitionRate will keep growing with block increase.
        // Each 5 blocks, the sortitionRate will increase 1%
        rateGrowStep = 5;
        firstVoteCost = 10; //ToDo: 
        // After the vote people add to 5, the vote will go into revealPhase.
        voterSelectedNum = 5;
        
        status = VoteStatus.created;
        //the qv method 
        qvMethod = QVMethod.Scheme1;
    }

  
    /**
     * description: add a candidate to this vote. Note that it would be preferable
                    in some situations to require candidates to be added only during the
                    Created phase to guarantee every voter could see candidates
                    before committing their vote.
     * params:
     * return:
     **/
    function addCandidate(bytes32 _name, address _account)
        private
        stopInEmergency
        returns (uint)
    {
        require(candidateCount < 255, "Can not add more than 255 candidates.");
        uint8 newCandidateCount = candidateCount;
        newCandidateCount = newCandidateCount.add(1);
        candidates[newCandidateCount].name = _name;
        candidates[newCandidateCount].account = _account;
        candidateCount = newCandidateCount;
        // Add candidate to Sortition module.
        addParticapator(_account);
    }


    function voteCommitPhase()
        public 
        onlyOwner()
    {
        status = VoteStatus.commitPhase;
    }

    /// @dev A method returns whether the user has the right to vote, the cost of votes.
    /// @param _account The User's account is used to participate in the sortition to determine if there is a chance to vote.
    /// @param _registerTime The User's Register time which must be eariler than the poll create time.
    /// @param _honor Weighted score of user attribute, which is used to magnify the chance of the sortition. Between 0 ~ 90%
    /// @param _numVotes The User's Votes.
    /// @returns _isChoice Returns whether it was selected by the sortition.  
    /// @returns _voteCost Returns the cost of the corresponding votes.
    /// voteCost = numVotes**2 * firstVoteCost
    function commitVoteAux(
        address _account,
        uint _registerTime,
        uint _honor,
        uint _numVotes
    )
        public
        view
        onlyOwner
    returns(
        bool _isChoice,
        uint _voteCost
    )
    {
        if(_registerTime > startTime){
            return (false, 0);
        }
        /// @notice: rN is larger than randommod
        uint rN = getRandomNum();
        /// @dev
        bytes32 hashWithRandom = keccak256(abi.encodePacked(_account,rN));
        uint res = uint(hashWithRandom) % rN;

        /// sortitionRate will increase with the block growing
        uint blockGrow = block.number - createBlock;
        uint bound = rN * (_honor + sortitionRate + blockGrow / rateGrowStep) / 100 ;

        if(res < bound){
            uint voteCost = calculateCost(_numVotes);
            return (true, voteCost);
        }else{
            return (false, 0);
        }

    }

    /// @dev An internal method for user to calculate the cost of vote
    function calculateCost(uint _voteNum)
        internal
        view
    returns(uint)
    {
        uint voteCost = (_voteNum.mul(_voteNum)).mul(firstVoteCost);
        return voteCost;
    }
    
    /// @dev The user votes, which can only be called by the contract creator.
    /// The method is necessary to prevent a sybil attack using many addresses
    /// to place votes at firstVoteCost.
    /// onlyOwner means this method can only be called by the contruct deployer
    function commitVote(address _voter, bytes32 commitment, uint32 _numVotes)
        public
        onlyOwner
        commitPhase
        stopInEmergency
    {
        approvedVoters[_voter] = true;
        /* Can submit fractions of votes, but must submit at least one
        vote, because voteRevealerRefund will return at least the cost
        of one vote */
        require (_numVotes >= 1, "Must submit at least one vote.");

        /* Can vote only one time per election (but can cast multiple
        votes at that time) */
        require (
            votes[_voter].numVotes == 0,
            "Can only commit votes once."
        );

        // Commit the vote(s) and adjust totalVotesCommitted and totalWeiPaid
        Vote memory newVote = Vote(uint32(_numVotes), commitment);
        votes[_voter] = newVote;
        
        voterCount = voterCount.add(1);
        if( voterCount >= voterSelectedNum)
        {
            voterCount = 0;
            status = VoteStatus.revealPhase;
        }
            
    }



    /**
     * description:Reveal votes.
     * params:
     * return:
     **/
    function revealVote(address _voter, uint _vote, bytes32 _salt)
        public
        onlyOwner
        stopInEmergency
        onlyApprovedVoters
    {
        // Check that hash(vote, numVotes, salt) == commitment
        Vote memory committedVote = votes[_voter];
        uint32 numVotes = committedVote.numVotes;
        bytes32 revealHash = keccak256(
            abi.encodePacked(
                uint(numVotes),
                _vote,
                _salt
            )
        );
        require (
            revealHash == committedVote.commitment,
            "Hash does not match the committed hash."
        );

        /* Delete the original vote to reduce gas costs and prevent the
        same commitment from being revealed multiple times. User's are
        still prevented from committing more votes by poll phase
        restrictions */
        delete votes[_voter];

        /* Count votes. Tally numVotes to the voteCount of the selected
        candidate */
        require (
            _vote > 0 && _vote <= candidateCount,
            "Vote must be placed for a valid candidate."
        );
        uint _voteCount = uint(candidates[_vote].voteCount);
        _voteCount = _voteCount.add(numVotes);
        candidates[_vote].voteCount = uint32(_voteCount);
        
        voterToCandidate[_voter] = _vote;
        
        voterCount = voterCount.add(1);
        if( voterCount >= voterSelectedNum)
        {
            voterCount = 0;
            status = VoteStatus.completed;
        }
           
    }
    

    function completeVote()
        public
        completePhase
        view
        returns(uint)
    {
        uint topicChoice = 0;
        for(uint i = 1; i < candidateCount; i++) {
            if
            (
                candidates[i].voteCount >
                candidates[topicChoice].voteCount
            )
            {
                topicChoice = i;
            }
        }
        return topicChoice;
    }

    /// @dev check the vote phase is completed
    function isVoteCompleted()
        public
        view
    returns(bool)
    {
        if(status == VoteStatus.completed)
            return true;
        return false;
    }

    /// @dev check the address is a voter and check is support for sucesss side
    function getVoterResult(address _voter)
        public
        view
    returns(bool isVoter, bool isWinSide, uint winChoice)
    {
        isVoter = approvedVoters[_voter];
        winChoice = completeVote();
        if(winChoice == voterToCandidate[_voter])
            isWinSide = true;
        else
            isWinSide = false;
    }
    
    
    function toggleContractActive() public {
        stopped = !stopped;
    }


    /** IMPORTANT: Hashing your vote on the blockchain is extremely
    insecure and completely defeats the purpose of a commit-reveal
    system, as everyone can see the vote you hashed. This function is
    included here for development to give anyone testing the contract
    an easy way to confirm that their locally-created hash matches
    the on-chain keccak256 used in the revealVote() function. */
    function hash(uint _numVotes, uint _vote, bytes32 salt)
    public view returns (bytes32) {
        /* For testing purposes, check that commitment vote is within
        the number of candidates in the poll to prevent confusion
        from committing an invalid vote */
        require (
            _vote <= candidateCount,
            "Vote must be placed for a valid candidate."
        );
        return keccak256(abi.encodePacked( _numVotes, _vote, salt));
    }
}
