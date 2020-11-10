import React from 'react'
import './datatable.styles.css'
import DataRow from '../../components/DataRow/DataRow.component'
import {Link} from 'react-router-dom'

class Datatable extends React.Component{
    constructor() {
        super();
        this.state = {
            fieldName: "everyFieldName",
            filterValue: "",
            openData: []
        }
    }

    componentDidMount(){
        fetch('http://localhost:8080/data', {
            credentials: "include",
            method: "GET"
        }).then(response => {
            return response.json()
        }).then(data=> {
            this.setState({
                openData: data.OpenData
            })
        })
    }

    handleFilterValueChange = event => {
        this.setState({
            filterValue: event.target.value
        })
    }
    handleFieldNameChange = event => {
        this.setState({
            fieldName: event.currentTarget.value
        })
    }

    handleExportFilteredDataInJSON = () => {
        fetch("http://localhost:8080/filteredJSON?fieldName=" + this.state.fieldName + "&&filterValue=" + this.state.filterValue, {
            credentials: "include",
            method: "GET"
        }).then((response) => response.blob())
            .then((blob) => {
                const url = window.URL.createObjectURL(new Blob([blob]));
                const link = document.createElement('a');
                link.href = url;
                link.setAttribute('download', `cars.json`);
                document.body.appendChild(link);
                link.click();
                link.parentNode.removeChild(link);
            })
    }
    handleExportFilteredDataInCSV = () => {
        fetch("http://localhost:8080/filteredCSV?fieldName=" + this.state.fieldName + "&&filterValue=" + this.state.filterValue, {
            credentials: "include",
            method: "GET"
        }).then((response) => response.blob())
            .then((blob) => {
                const url = window.URL.createObjectURL(new Blob([blob]));
                const link = document.createElement('a');
                link.href = url;
                link.setAttribute('download', `cars.csv`);
                document.body.appendChild(link);
                link.click();
                link.parentNode.removeChild(link);
            })
    }

    handleFilterDataButtonClick = () => {
        fetch("http://localhost:8080/filteredData?fieldName=" + this.state.fieldName + "&&filterValue=" + this.state.filterValue, {
            credentials: "include",
            method: "GET"
        }).then(response => {
            return response.json()
        }).then(data=> {
            this.setState({
                openData: data.OpenData
            })
        })
    }

    render(){
        return(
            <div className='page-container'>
                <div className='filter-container'>
                    <input type='text' onChange={this.handleFilterValueChange} value={this.state.filterValue}/>
                    <select onChange={this.handleFieldNameChange} defaultValue='everyFieldName'>
                        <option value="everyFieldName">Sva polja (wildcard)</option>
                        <option value="vehicleBrand">vehicleBrand</option>
                        <option value="model">model</option>
                        <option value="year">year</option>
                        <option value="engineName">engineName</option>
                        <option value="engineType">engineType</option>
                        <option value="powerOutput">powerOutput</option>
                        <option value="torque">torque</option>
                        <option value="topSpeed">topSpeed</option>
                        <option value="acceleration">acceleration</option>
                        <option value="weight">weight</option>
                        <option value="bootSpace">bootSpace</option>
                        <option value="awardEvent">awardEvent</option>
                        <option value="award">award</option>
                    </select>
                    <button onClick={this.handleFilterDataButtonClick}>Pretrazi</button>
                    <button onClick={this.handleExportFilteredDataInJSON}>Export u JSON</button>
                    <button onClick={this.handleExportFilteredDataInCSV}>Export u CSV</button>
                    <Link to= '/'>Index</Link>
                </div>
                <div className='table-container'>
                   <div className='table-header'>
                       <p>Brand</p>
                       <p>Year</p>
                       <p>Year</p>
                       <p>Engine name</p>
                       <p>Engine type</p>
                       <p>Power Output(PS)</p>
                       <p>Torque(N/m)</p>
                       <p>Top speed(km/h)</p>
                       <p>Acceleration(0-100km/h)</p>
                       <p>Weight(kg)</p>
                       <p>Boot space(L)</p>
                       <p>Wiki-sufix</p>
                       <p>Awards</p>
                   </div>
                    {this.state.openData.map(car => (
                        <DataRow  key = { car.Model} acceleration = {car.Acceleration} awards = {car.Awards} bootSpace = {car.BootSpace} engineType = {car.Engine.Type} engineName = {car.Engine.Name} model = {car.Model}
                                  powerOutput = {car.PowerOutput} topSpeed = {car.TopSpeed} torque = {car.Torque} weight = {car.Weight} WikipediaSufix = {car.WikipediaSufix} year = {car.Year} vehicleBrand = {car.VehicleBrand}/>
                    ))
                    }
                </div>
            </div>
        )
    }
}

export default Datatable