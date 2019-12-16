pragma solidity ^0.4.21;
pragma experimental ABIEncoderV2;
import "./Struct_definition.sol";

contract dream_factory is myStruct{
address master;
Project[] projects;

//Member[] members;
mapping (address => Member) members;

Invest[] invests;
uint num_invest;
event updated_name(string project_name, string new_projcet_name);
event updated_goal(Goal project_goal, Goal new_project_goal);

function dream_factory() public {
    master = msg.sender;
    num_invest = 0;
}

function get_proposer (uint _project_id) private view returns (Member) {
    return projects[_project_id].proposer;
}

function get_investor_list (uint _project_id) private view returns (Member[]) {
    return projects[_project_id].investor_list;
}

function get_investor_count (uint _project_id) public view returns (uint) {
    return projects[_project_id].contribute_count;
}

function get_date(uint _project_id) public view returns (DateTime) {
    return projects[_project_id].date;
}

function get_member(address _member_id) private returns (Member) {
    return members[_member_id];
}

function create_member(uint _db_index) public returns(bool) {
    Invest[] tmp;
    Project[] tempj;
    members[msg.sender] = Member(_db_index,tmp,tempj,msg.sender);
    return true;
}

function validate_check(uint _project_id, bool _milestone, bool _scam_check, bool _investor_check) public view returns (bool) {
    projects[_project_id].project_goal.milestone = _milestone;
    projects[_project_id].project_goal.scam_check = _scam_check;
    projects[_project_id].project_goal.investor_check = _investor_check;
    if(_milestone && _scam_check  && _investor_check)
        return true;
    else
        return false;
}
function withdraw(uint _project_id, uint _amount, bool _milestone, bool _scam_check, bool _investor_check) payable public returns (bool) {
    if(validate_check(_project_id, _milestone, _scam_check, _investor_check)) {
        projects[_project_id].proposer.wallet.transfer(_amount);
        return true;
    }
    else
        return false;
}
function invest(uint _project_id, uint _amount, uint _target_project, uint _value) public {
    projects[_project_id].contributing[msg.sender] = msg.value;
    assert(_amount != msg.value);
    projects[_project_id].investor_list.push(members[msg.sender]);
    invests.push(Invest(num_invest++, msg.sender, _target_project, msg.value));
}
function delete_project(uint _project_id) public returns (bool success) {
    delete projects[_project_id];
    return true;
}
function update_project(uint _project_id, string new_project_name, Goal new_project_goal) public returns (bool success) {
    require(msg.sender == projects[_project_id].proposer.wallet);
    projects[_project_id].project_name = new_project_name;
    projects[_project_id].project_goal = new_project_goal;
    emit updated_name(projects[_project_id].project_name, new_project_name);
    emit updated_goal(projects[_project_id].project_goal, new_project_goal);
    return true;
}
}