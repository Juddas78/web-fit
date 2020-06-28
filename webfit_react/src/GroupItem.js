import React from 'react';

class GroupItem extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
          isToggleOn: true,
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
        const { name, isToggleOn } = this.state;
        return (
            <button onClick={this.handleClick}>
            {name}: {isToggleOn ? 'ON' : 'OFF'}
            </button>
      );
    }
  }

  export default GroupItem;
