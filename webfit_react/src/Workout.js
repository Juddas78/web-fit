import React from 'react';
import './Workout.css';


class MyComponent extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        error: null,
        isLoaded: false,
        exercises: []
      };
    }
  
    componentDidMount() {
      fetch("/workout")
        .then((res) => res.json())
        .then(
            (result) => {
            console.log(result);
            this.setState({
              isLoaded: true,
              exercises: result.exercises
            });
          },
          // Note: it's important to handle errors here
          // instead of a catch() block so that we don't swallow
          // exceptions from actual bugs in components.
          (error) => {
            this.setState({
              isLoaded: true,
              error
            });
          }
        )
    }
  
    render() {
      const { error, isLoaded, exercises } = this.state;
      if (error) {
        return <div>Error: {error.message}</div>;
      } else if (!isLoaded) {
        return <div>Loading...</div>;
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
  export default MyComponent;