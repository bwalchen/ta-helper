import React, { Component } from 'react'; //import React Component
import Time from 'react-time'
import './styles/Questions.css';
import QuestionModal from './QuestionModal';
import {Button} from 'reactstrap'


export default class QuestionList extends Component {
  constructor(props){
    super(props);
    this.state = {};
  }

  componentWillMount(){
    console.log(this.props.questions)
  }

  render() {
    let questionItems = [];
    let questions = this.props.questions;
    if(questions){
      let keyArray = Object.keys(questions).map(key => {
        questions[key].dbID = questions[key].id 
        questions[key].id = key;
        return questions[key];
      });
      keyArray.sort((a,b) => a.time - b.time);
      questionItems = keyArray.map( each =>  <QuestionItem  key={each.id} 
        deletequestionCallback={this.props.deletequestionCallback} 
        editquestionCallback={this.props.editquestionCallback} question={each} 
        currentUser={this.props.currentUser}/>)
    }
    return (
      <div aria-live="polite" className="container">
          {questionItems}
      </div>);
  }
}


class QuestionItem extends Component {
  deletequestionHandler(){
    this.props.deletequestionCallback(this.props.question.dbID)
  }
  render() {
    let question = this.props.question;
    let user = JSON.parse(localStorage.getItem('User'));
    console.log(question)
    return (
      <div className="row py-4 bg-white border">
        <div className="col pl-4 pl-lg-1">
          <img src='./../imgs/up-arrow.png'></img>
          <img src='./../imgs/down-arrow.png'></img>
          <div className="question">{question.questBody}</div>
        </div>
        {this.props.user.role === 'instructor' ? 
        '' : ''} 
        {/* {this.props.question.creator.id ===  user.id ?
        <span>
        <questionModal question={question} buttonCallback={this.props.editquestionCallback}></questionModal>
        <Button color="danger" className="float-right" onClick={(e) => this.deletequestionHandler()}>Delete</Button>
        </span> : ''} */}}
      </div>      
    );
  }
}
