import React from 'react'
import './index.styles.css'
import {Link} from 'react-router-dom'

const Index = () => (
    <div>
        <div className='group'>
            <button onClick={handleExportInJSONButtonClick}>Export in JSON</button>
            <button onClick={handleExportInCSVButtonClick}>Export in CSV</button>
            <Link to='/datatable'>Datatable</Link>
        </div>
        <div className='data-description'>
            <p>Autor: Eduard Đuras</p>
            <p> Licenca: Creative commons Attribution-NonCommercial 4.0 International (CC BY-NC 4.0)</p>
            <p>Verzija: 1.0</p>
            <p>Naziv: Marke, modeli i specifikacije automobila</p>
            <p> Jezik podataka: engleski</p>
            <p>Formati: JSON i CSV</p>
            <p> Ključne riječi: auto, motor, specifikacije, nagrade</p>
            <p>Datum objave: 26.10.2020</p>
            <hr/>
            <h2>Opis podataka</h2>
            <p className='dataset-description'>Svaki zapis sadrži isti skup podataka. Jedan zapis predstavlja jedan model automobila. Razlikujemo modele automobila koji su izašli različitih godina. Svaki zapis sadrži i podatke o motoru koji se nalazi u automobilu, njegovo ime, kojeg je tipa (benzin ili dizel, proširivo u budućnosti na hibridne i električne motore) te koje su mogućnosti motora (snaga, najveća brzina, akceleracija...). Prate se i podaci o težini vozila te veličini prtljažnika. Dodani su podaci o nagradama koje je automobil osvojio, npr. nagrada od nekog časopisa, emisije...
            </p>
        </div>
    </div>
);

function handleExportInJSONButtonClick(){
    fetch('http://localhost:8080/JSON', {
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
function handleExportInCSVButtonClick(){
    fetch('http://localhost:8080/CSV', {
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

export default Index