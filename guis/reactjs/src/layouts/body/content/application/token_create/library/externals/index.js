import React from 'react';
import Table from 'react-bootstrap/Table';
import Stack from 'react-bootstrap/Stack';
import Button from 'react-bootstrap/Button';

export class ExternalGrammars extends React.Component {
  render() {
    return (
        <div>
            <Table striped bordered hover size="sm">
              <thead>
                <tr>
                    <th>Name</th>
                    <th>Grammar</th>
                    <th>Action</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                <td><a href="/">myLocalName</a></td>
                <td><a href="/">myGrammar</a></td>
                  <td>
                      <Stack className="justify-content-center" direction="horizontal" gap={1}>
                          <Button variant="success" size="sm"><i className="bi bi-pencil"></i></Button>
                          <Button variant="danger" size="sm"><i className="bi bi-trash"></i></Button>
                      </Stack>
                  </td>
                </tr>
              </tbody>
            </Table>
            <Button variant="primary" className="add bi bi-folder-plus"><span>New External Grammar</span></Button>
      </div>
    );
  }
}
