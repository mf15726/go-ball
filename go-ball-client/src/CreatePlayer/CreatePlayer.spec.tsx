import React, { useState } from 'react';
import { CreatePlayer } from './CreatePlayer';
import { shallow, configure } from 'enzyme';
import fetchMock from 'fetch-mock';
import { urlBase } from "../../constsAndTypes/consts";
import { Button } from '@material-ui/core';
import Adapter from 'enzyme-adapter-react-16';

describe("CreatePlayer Component", () => {

    configure({ adapter: new Adapter() })

    const component = shallow(<CreatePlayer playerMade={() => (setPlayerCreated(true))}/>)

    it('should match snapshot', () => {
        expect(component).toMatchSnapshot();
    });

    it('should return a 200 whenever the button is pressed', () => {
        let playerName = undefined
        let stackSize = undefined
        let tablePosition = undefined
        fetchMock.post(urlBase + "/player/makePlayer?playerName=" + playerName + "&stackSize=" + stackSize +
        "&tablePosition=" + tablePosition, 200)
        expect(component.find(Button).first().dive().simulate('click'))
    })

});