import React, { Component } from 'react';
import { connect } from 'react-redux';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import {
  shortenURL,
} from '../../redux/actions/mainFormActions';

class MainForm extends Component {
  constructor(props) {
    super();

    this.state = {
      longURL: '',
    }

    this.handleLongURLChange = this.handleLongURLChange.bind(this);
    this.handleShortenURL = this.handleShortenURL.bind(this);
  }

  handleLongURLChange(e) {
    this.setState({ longURL: e.target.value });
  }

  handleShortenURL() {
    const {
      shortenURL,
    } = this.props;
    
    shortenURL(this.state.longURL);
    
    this.setState({
      longURL: '',
    });
  }

  render() {
    return (
      <div className="form-wrapper">
        <Form>
          <Form.Group as={Row} controlId="exampleForm.ControlTextarea1">
            <Col>
              <Form.Control as="textarea" rows="3" placeholder="Ваша длинная ссылка" onChange={this.handleLongURLChange}/>
            </Col>
          </Form.Group>
          <Form.Group as={Row}>
            <Button onClick={this.handleShortenURL} variant="primary" type="submit">Укоротить</Button>
          </Form.Group>
        </Form>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return { 
    longURL: state.mainForm.longURL,
  };
}

function mapDispatchToProps(dispatch) {
  return{
    shortenURL: bindActionCreators(shortenURL, dispatch),
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(MainForm);
