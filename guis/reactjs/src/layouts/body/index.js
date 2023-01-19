import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

import { LeftMenu } from './leftmenu/index.js';
import { Content } from './content/index.js';

import './index.css';

export class Body extends React.Component {
  render() {
    return (
        <div id="body">
            <Container fluid>
                <Row>
                    <Col md="auto"><LeftMenu /></Col>
                    <Col><Content /></Col>
                </Row>
            </Container>
        </div>
    );
  }
}
