import React from 'react';
import Workout from './Workout'
import Groups from './GroupsForm'

class CreateWorkout extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            exercises: [],
            groups: ["chest", "upper back", "middle back", "shoulders", "biceps", "triceps", "lower back", "glutes", "quads", "hamstrings", "calves", "abs"],
            userGroup: []
          };
        this.handleClick = this.handleClick.bind(this);
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

      handleClick() {
        console.log(this.state.userGroup)
        fetch("/workout", {
          method: 'POST', // *GET, POST, PUT, DELETE, etc.
          mode: 'cors', // no-cors, *cors, same-origin
          cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
          credentials: 'same-origin', // include, *same-origin, omit
          headers: {
            'Content-Type': 'application/json'
            // 'Content-Type': 'application/x-www-form-urlencoded',
          },
          redirect: 'follow', // manual, *follow, error
          referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
          body: this.state.userGroup // body data type must match "Content-Type" header
        })
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
        const { exercises, userGroup, groups } = this.state;
        return (
            <div>
                <h1>Create Workout</h1>
                <Workout exercises={exercises}/>
                <Groups userGroup={ userGroup } groups={groups} />
                <button onClick={this.handleClick}>
                  Get Workout
                </button>
            </div>
    );
    }
}
export default CreateWorkout;
