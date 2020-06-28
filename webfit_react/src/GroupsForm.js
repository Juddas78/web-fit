import React from 'react';
import GroupItem from './GroupItem'
import './GroupsForm.css'
class Groups extends React.Component {
    inputGroups = []
    constructor(props) {
      super(props);
      this.state = {
          isToggleOn: true,
        };
  
      // This binding is necessary to make `this` work in the callback
      this.selectGroup = this.selectGroup.bind(this);
      this.inputGroups = this.props.userGroup;
    }

    selectGroup(item) {
        const index = this.props.userGroup.indexOf(item);
        if( index > -1) {
            this.props.userGroup.splice(index, 1)
            console.log(item, ' removed!', this.props.userGroup);
        } else {
            this.props.userGroup.push(item)
            console.log(item, ' added!', this.props.userGroup);
        }
    }
  
    render() {
    // const { muscleGroups } = this.state;
    const { groups } = this.props;
      return (
        <div className="groups-div">
            {groups.map(group => (
                <GroupItem name={group} isToggleOn='OFF' selectGroup={this.selectGroup}></GroupItem>
            ))}
        </div>
      );
    }
  }

  export default Groups;
