import React from 'react';
import Table from 'react-bootstrap/Table'
import axios from 'axios';

class TenantList extends React.Component {

    readData() {
        const self = this;
        axios.get(window.global.api_location+'/tenants').then(function(response) {
            console.log(response.data);

            self.setState({tenants: response.data});
        }).catch(function (error){
            console.log(error);
        });
    }

    getTenants() {
        let table = []

        for (let i=0; i < this.state.tenants.length; i++) {

            table.push(
            <tr key={i}>
                <td>{this.state.tenants[i].id}</td>
                <td>{this.state.tenants[i].name}</td>
                <td>{this.state.tenants[i].description}</td>
                <td>{this.state.tenants[i].kbo}</td>
            </tr>
            );
        }

        return table
    }

    constructor(props) {
        super(props);
        this.readData();
        this.state = {tenants: []};
    
        this.readData = this.readData.bind(this);
    }

    render() {
      return (
        <div>
            <h1 style={{marginBottom: "40px"}}>Tenants</h1>
            <Table>
                <thead>
                    <tr>
                        <th>
                            ID
                        </th>
                        <th>
                            Name
                        </th>
                        <th>
                            Description
                        </th>
                        <th>
                            KBO
                        </th>
                    </tr>
                </thead>
                <tbody>
                    {this.getTenants()}
                </tbody>
            </Table>
        </div>
      ) 
    }
}

export default TenantList;