import React from 'react';
import './App.css';
import { Switch, Route } from 'react-router-dom';
import Index from '../src/pages/index/index.component'
import Datatable from './pages/dataTable/datatable.component';

class App extends React.Component {


  render(){
    return (
        <div className="App">
          <div className = 'page-container'>
            <Switch>
              <Route exact path="/" component={Index} />
              <Route exact path="/datatable" component={Datatable} />
            </Switch>
          </div>
        </div>
    )
  }
}

export default App;
