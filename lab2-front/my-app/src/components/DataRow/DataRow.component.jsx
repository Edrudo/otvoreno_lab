import React from 'react'
import  './DataRow.styles.css'
import Awards from '../awards/awards.component'

const DataRow = ({acceleration, bootSpace, engineType, engineName, model, powerOutput, topSpeed, torque, vehicleBrand, weight, WikipediaSufix, year, awards}) => (
    <div className='row'>
            <p>{vehicleBrand}</p>
            <p>{model}</p>
            <p>{year}</p>
            <p>{engineName}</p>
            <p>{engineType}</p>
            <p>{powerOutput}</p>
            <p>{torque}</p>
            <p>{topSpeed}</p>
            <p>{acceleration}</p>
            <p>{weight}</p>
            <p>{bootSpace}</p>
            <p>{WikipediaSufix}</p>
            <Awards awards = {awards}/>
    </div>
);

export default DataRow