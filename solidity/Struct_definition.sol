pragma solidity ^0.4.24;


contract myStruct{ 



struct Goal {
    bool milestone;
    bool scam_check;
    bool investor_check;
}

struct Invest {
    uint invest_id;
    address investor;
    uint target_project;
    uint value;
}

struct Member {
    uint member_id;
    address wallet;
}

struct Project {
    Member[] investor_list;
    uint project_id;
    uint balance;
    Member proposer;
    uint contribute_count;
    uint date;
    bool validate;
    string project_name;
    Goal project_goal;
    mapping(address => uint) contributing;
}
}