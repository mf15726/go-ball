import React from 'react';
import { Button, TextField } from '@material-ui/core';
import { urlBase } from '../util/consts';
import './CreatePlayer.css';

interface CreatePlayerProps {
    playerMade : (playerCreated: any) => void
}

export function CreatePlayer(props: CreatePlayerProps) {

    let playerName: string;
    
    const makePlayer = () => {
        return fetch(urlBase + "/player/makePlayer?playerName=" + playerName, {
             mode: 'no-cors',
             method: 'POST',
             headers: {
                 'Content-Type': 'application/json',                
             }
         })
         .then(() => props.playerMade(true))
         .catch(_err => console.log("Error Making Player: " + _err));
    }

    const renderTitle = () => {
        return(
            <div className='CreatePlayerTitle'>
                Create Player
            </div>
        )
    }

    const renderTextFields = () => {
        return(
            <div className='textFieldContainer'>
                <div className = 'textField'>
                    <TextField 
                        required
                        id="player-name"
                        label="Player Name"
                        variant="outlined"
                        onChange={(e: any) => playerName = e.target.value} 
                    />
                </div>
                <div className = 'textField'>
                    <TextField
                        required
                        id="stack-size"
                        label="Stack Size"
                        type="number"
                        variant="outlined"
                        InputLabelProps={{
                        shrink: true,
                        }}
                        onChange={(e: any) => stackSize = e.target.value}
                    />
                </div>
                <div className = 'textField'>
                    <TextField
                        required
                        id="table-position"
                        label="Table Position"
                        type="number"
                        variant="outlined"
                        InputLabelProps={{
                        shrink: true,
                        }}
                        onChange={(e: any) => tablePosition = e.target.value}
                    />
                </div>
            </div>
        )
    }

    const renderButton = () => {
        return(
            <Button
                variant='contained'
                color='primary'
                onClick={() => makePlayer()} >
                Join Game
            </Button>
        )
    }

    return(
        <div className='CreatePlayerContainer'>
            {renderTitle()}
            {renderTextFields()}
            {renderButton()}
        </div>
    )

}