import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Badge from 'react-bootstrap/Badge';
import Button from 'react-bootstrap/Button';
import Tab from 'react-bootstrap/Tab';
import ListGroup from 'react-bootstrap/ListGroup';

import { Library } from './library/index.js';

import './index.css';


class Possibilities extends React.Component {
  render() {
    return (
        <div id="possibilities">
        </div>
    );
  }
}

class Grammar extends React.Component {
  render() {
    return (
        <div id="grammar">
            <h1>My Grammar</h1>
            <h2><Badge bg="success">Root</Badge> <a href="/">myRootToken</a></h2>
            <div class="token">
                <h3>Tokens</h3>
                <p>
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce eget rutrum arcu.
                    Nullam vel molestie nulla, sit amet feugiat velit. Nam sed ipsum molestie, tempor risus quis, lobortis risus.
                    Etiam arcu augue, euismod et nunc ac, aliquam fermentum elit. Nunc sodales dui vitae maximus finibus.
                    Integer risus velit, euismod non mi id, tincidunt condimentum lorem. Nam id odio in felis euismod accumsan.
                    Aenean ut fermentum velit, ut tristique tellus.
                </p>

                <div id="tokenslist">
                    <Tab.Container id="list-group-tabs-example" defaultActiveKey="#link1">
                        <Row>
                            <Col sm={4}>
                                <ListGroup as="ol" numbered>
                                    <ListGroup.Item
                                        className="d-flex justify-content-between align-items-start"
                                        action
                                        href="#link1"
                                    >
                                        <div className="ms-2 me-auto">
                                            <div className="fw-bold">firstToken</div>
                                        </div>
                                        <Badge bg="primary" pill>
                                            32
                                        </Badge>
                                    </ListGroup.Item>
                                    <ListGroup.Item
                                        className="d-flex justify-content-between align-items-start"
                                        action
                                        href="#link2"
                                    >
                                        <div className="ms-2 me-auto">
                                            <div className="fw-bold">secondToken</div>
                                        </div>
                                        <Badge bg="primary" pill>
                                            14
                                        </Badge>
                                    </ListGroup.Item>
                                    <ListGroup.Item
                                        className="d-flex justify-content-between align-items-start"
                                        action
                                        href="#link3"
                                    >
                                        <div className="ms-2 me-auto">
                                            <div className="fw-bold">thirdToken</div>
                                        </div>
                                        <Badge bg="primary" pill>
                                            22
                                        </Badge>
                                    </ListGroup.Item>
                                </ListGroup>
                            </Col>
                            <Col sm={8}>
                                <Tab.Content>
                                    <Tab.Pane eventKey="#link1">
                                        <Possibilities />
                                    </Tab.Pane>
                                    <Tab.Pane eventKey="#link2">
                                        <Possibilities />
                                    </Tab.Pane>
                                    <Tab.Pane eventKey="#link3">
                                        <Possibilities />
                                    </Tab.Pane>
                                </Tab.Content>
                            </Col>
                        </Row>
                    </Tab.Container>
                </div>
            </div>
        </div>
    );
  }
}

class MainPanel extends React.Component {
  render() {
    return (
        <div id="mainpanel">
            <Grammar />
        </div>
    );
  }
}

export class TokenCreate extends React.Component {
  render() {
    return (
        <Container>
            <Row>
                <Col><MainPanel /></Col>
                <Col md="auto"><Library /></Col>
            </Row>
        </Container>
    );
  }
}
