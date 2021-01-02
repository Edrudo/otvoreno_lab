import React from 'react'
import './awards.styles.css'

var i = 0

const Awards = ({awards}) => (
    <div className='awards-container'>
        {awards.map(award => (<p key={i++}>"{award.Award}" by {award.Event}</p>))}
    </div>
);

export default Awards