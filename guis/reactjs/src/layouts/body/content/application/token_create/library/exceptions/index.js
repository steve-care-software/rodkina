import React from 'react';
import Table from 'react-bootstrap/Table';
import Stack from 'react-bootstrap/Stack';
import Button from 'react-bootstrap/Button';

export class Exceptions extends React.Component {
  render() {
    return (
        <div id="exceptions">
            <Table striped bordered hover size="sm">
              <thead>
                <tr>
                    <th>Name</th>
                    <th>Exception</th>
                    <th>Escaspe</th>
                    <th>Action</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                <td><a href="/">myName</a></td>
                <td><a href="/">myException</a></td>
                <td><a href="/">myEscape</a></td>
                  <td>
                      <Stack className="justify-content-center" direction="horizontal" gap={1}>
                          <Button variant="success" size="sm"><i className="bi bi-pencil"></i></Button>
                          <Button variant="danger" size="sm"><i className="bi bi-trash"></i></Button>
                      </Stack>
                  </td>
                </tr>
                <tr>
                <td><a href="/">myOtherName</a></td>
                <td><a href="/">otherExp</a></td>
                <td>-</td>
                    <td>
                          <Stack className="justify-content-center" direction="horizontal" gap={1}>
                              <Button variant="success" size="sm"><i className="bi bi-pencil"></i></Button>
                              <Button variant="danger" size="sm"><i className="bi bi-trash"></i></Button>
                          </Stack>
                    </td>
                </tr>
              </tbody>
            </Table>
            <Button variant="primary" className="add bi bi-folder-plus"><span>New Exception</span></Button>
      </div>
    );
  }
}
