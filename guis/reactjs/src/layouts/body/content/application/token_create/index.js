import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

import { Library } from './library/index.js';

class MainPanel extends React.Component {
  render() {
    return (
        <div id="mainpanel">
            Main panel
        </div>
    );
  }
}

export class TokenCreate extends React.Component {
  render() {
    return (
        <Container fluid>
            <Row>
                <Col><MainPanel /></Col>
                <Col md="auto"><Library /></Col>
            </Row>
        </Container>
    );
  }
}
