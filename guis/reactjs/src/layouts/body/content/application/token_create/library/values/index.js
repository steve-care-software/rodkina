import React from 'react';
import Table from 'react-bootstrap/Table';
import Stack from 'react-bootstrap/Stack';
import Button from 'react-bootstrap/Button';

export class Values extends React.Component {
  render() {
    return (
        <div>
            <Table striped bordered hover size="sm">
              <thead>
                <tr>
                  <th>Variable</th>
                  <th>Decimal</th>
                  <th>Symbol</th>
                  <th>Description</th>
                  <th>Action</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>myValue</td>
                  <td>21</td>
                  <td>NAK</td>
                  <td>Negative Acknowledge</td>
                  <td>
                      <Stack className="justify-content-center" direction="horizontal" gap={1}>
                          <Button variant="success" size="sm"><i className="bi bi-pencil"></i></Button>
                          <Button variant="danger" size="sm"><i className="bi bi-trash"></i></Button>
                      </Stack>
                  </td>
                </tr>
                <tr>
                  <td>myOther</td>
                  <td>48</td>
                  <td>0</td>
                  <td>Zero</td>
                  <td>
                      <Stack className="justify-content-center" direction="horizontal" gap={1}>
                          <Button variant="success" size="sm"><i className="bi bi-pencil"></i></Button>
                          <Button variant="danger" size="sm"><i className="bi bi-trash"></i></Button>
                      </Stack>
                  </td>
                </tr>
                <tr>
                  <td>myValue</td>
                  <td>21</td>
                  <td>NAK</td>
                  <td>Negative Acknowledge</td>
                  <td>
                      <Stack className="justify-content-center" direction="horizontal" gap={1}>
                          <Button variant="success" size="sm"><i className="bi bi-pencil"></i></Button>
                          <Button variant="danger" size="sm"><i className="bi bi-trash"></i></Button>
                      </Stack>
                  </td>
                </tr>
                <tr>
                  <td>myOther</td>
                  <td>48</td>
                  <td>0</td>
                  <td>Zero</td>
                  <td>
                      <Stack className="justify-content-center" direction="horizontal" gap={1}>
                          <Button variant="success" size="sm"><i className="bi bi-pencil"></i></Button>
                          <Button variant="danger" size="sm"><i className="bi bi-trash"></i></Button>
                      </Stack>
                  </td>
                </tr>
                <tr>
                  <td>myValue</td>
                  <td>21</td>
                  <td>NAK</td>
                  <td>Negative Acknowledge</td>
                  <td>
                      <Stack className="justify-content-center" direction="horizontal" gap={1}>
                          <Button variant="success" size="sm"><i className="bi bi-pencil"></i></Button>
                          <Button variant="danger" size="sm"><i className="bi bi-trash"></i></Button>
                      </Stack>
                  </td>
                </tr>
                <tr>
                  <td>myOther</td>
                  <td>48</td>
                  <td>0</td>
                  <td>Zero</td>
                  <td>
                      <Stack className="justify-content-center" direction="horizontal" gap={1}>
                          <Button variant="success" size="sm"><i className="bi bi-pencil"></i></Button>
                          <Button variant="danger" size="sm"><i className="bi bi-trash"></i></Button>
                      </Stack>
                  </td>
                </tr>
                <tr>
                  <td>myValue</td>
                  <td>21</td>
                  <td>NAK</td>
                  <td>Negative Acknowledge</td>
                  <td>
                      <Stack className="justify-content-center" direction="horizontal" gap={1}>
                          <Button variant="success" size="sm"><i className="bi bi-pencil"></i></Button>
                          <Button variant="danger" size="sm"><i className="bi bi-trash"></i></Button>
                      </Stack>
                  </td>
                </tr>
                <tr>
                  <td>myOther</td>
                  <td>48</td>
                  <td>0</td>
                  <td>Zero</td>
                  <td>
                      <Stack className="justify-content-center" direction="horizontal" gap={1}>
                          <Button variant="success" size="sm"><i className="bi bi-pencil"></i></Button>
                          <Button variant="danger" size="sm"><i className="bi bi-trash"></i></Button>
                      </Stack>
                  </td>
                </tr>
                <tr>
                  <td>myValue</td>
                  <td>21</td>
                  <td>NAK</td>
                  <td>Negative Acknowledge</td>
                  <td>
                      <Stack className="justify-content-center" direction="horizontal" gap={1}>
                          <Button variant="success" size="sm"><i className="bi bi-pencil"></i></Button>
                          <Button variant="danger" size="sm"><i className="bi bi-trash"></i></Button>
                      </Stack>
                  </td>
                </tr>
                <tr>
                  <td>myOther</td>
                  <td>48</td>
                  <td>0</td>
                  <td>Zero</td>
                  <td>
                      <Stack className="justify-content-center" direction="horizontal" gap={1}>
                          <Button variant="success" size="sm"><i className="bi bi-pencil"></i></Button>
                          <Button variant="danger" size="sm"><i className="bi bi-trash"></i></Button>
                      </Stack>
                  </td>
                </tr>
                <tr>
                  <td>myValue</td>
                  <td>21</td>
                  <td>NAK</td>
                  <td>Negative Acknowledge</td>
                  <td>
                      <Stack className="justify-content-center" direction="horizontal" gap={1}>
                          <Button variant="success" size="sm"><i className="bi bi-pencil"></i></Button>
                          <Button variant="danger" size="sm"><i className="bi bi-trash"></i></Button>
                      </Stack>
                  </td>
                </tr>
                <tr>
                  <td>myOther</td>
                  <td>48</td>
                  <td>0</td>
                  <td>Zero</td>
                  <td>
                      <Stack className="justify-content-center" direction="horizontal" gap={1}>
                          <Button variant="success" size="sm"><i className="bi bi-pencil"></i></Button>
                          <Button variant="danger" size="sm"><i className="bi bi-trash"></i></Button>
                      </Stack>
                  </td>
                </tr>
              </tbody>
            </Table>
            <Button variant="primary" className="add bi bi-folder-plus"><span>New Value</span></Button>
      </div>
    );
  }
}
