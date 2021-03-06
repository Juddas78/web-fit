import React from 'react';
import './ExerciseList.css';


class ExerciseList extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        error: null,
        isLoaded: props.exercises ? true : false,
      };
    }
    
    componentDidMount() {
      console.log(this.props.exercises);
      this.setState({
        isLoaded: true,
      });
    }
    
    handleClick() {
        this.setState(state => ({
          isToggleOn: !state.isToggleOn
        }));
    }

  
    render() {
      const { error, isLoaded } = this.state;
      const { exercises } = this.props
      if (error) {
        return <div>Error: {error.message}</div>;
      } else if (!isLoaded) {
        return <div>Loading...</div>;
      } else if (exercises === null) {
        return (
          <ul id="lifts">
              <li>
                <span className="lift-name">Enter some muscle groups to get your workout</span>
              </li>
          </ul>
        )
      } else {
        return (
          <ul id="lifts">
            {exercises.map((lift) => (
              <li key={lift.name}>
                <span className="lift-name">{lift.name}</span> {lift.sets.length} sets: {lift.sets.map(set => (<span className="sets">{set.reps},</span>) )}
              </li>
            ))}
          </ul>
        );
      }
    }
  }
  export default ExerciseList;