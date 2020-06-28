import React from 'react';
import './GroupItem.css';



class GroupItem extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
          isToggleOn: false,
          name: props.name
        };
  
      // This binding is necessary to make `this` work in the callback
      this.handleClick = this.handleClick.bind(this);
    }
  
    handleClick(e) {
        this.props.selectGroup(this.state.name)
        this.setState(state => ({
            isToggleOn: !state.isToggleOn
        }));

    }
  
    render() {
      let btn_class = "btn "
        const { name, isToggleOn} = this.state;
        return (
            <div className={isToggleOn ? "btn on" : "btn off"} onClick={this.handleClick}>
            <p>{name}</p>
            </div>
      );
    }
  }

  export default GroupItem;
