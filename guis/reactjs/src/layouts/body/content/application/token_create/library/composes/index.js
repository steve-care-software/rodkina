import React from 'react';
import Table from 'react-bootstrap/Table';
import Stack from 'react-bootstrap/Stack';
import Button from 'react-bootstrap/Button';

export class Composes extends React.Component {
  render() {
    return (
        <div id="channels">
            <Table striped bordered hover size="sm">
              <thead>
                <tr>
                  <th>Token</th>
                  <th>Previous</th>
                  <th>Next</th>
                  <th>Action</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td><a href="/">myToken</a></td>
                  <td><a href="/">myPrev</a></td>
                  <td><a href="/">myNext</a></td>
                  <td>
                      <Stack className="justify-content-center" direction="horizontal" gap={1}>
                          <Button variant="success" size="sm"><i className="bi bi-pencil"></i></Button>
                          <Button variant="danger" size="sm"><i className="bi bi-trash"></i></Button>
                      </Stack>
                  </td>
                </tr>
                <tr>
                    <td><a href="/">myToken</a></td>
                    <td>-</td>
                    <td><a href="/">myNext</a></td>
                    <td>
                          <Stack className="justify-content-center" direction="horizontal" gap={1}>
                              <Button variant="success" size="sm"><i className="bi bi-pencil"></i></Button>
                              <Button variant="danger" size="sm"><i className="bi bi-trash"></i></Button>
                          </Stack>
                    </td>
                </tr>
                <tr>
                    <td><a href="/">myToken</a></td>
                    <td><a href="/">myPrev</a></td>
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
            <Button variant="primary" className="add bi bi-folder-plus"><span>New Compose</span></Button>
      </div>
    );
  }
}
