import { expect } from 'chai';
import { describe, it, beforeEach, afterEach } from 'mocha';
import * as sinon from 'sinon';
import {
    STATE_OFF,
    riceAndWaterAdded,
    riceCookerState,
    displayState,
    plugIn,
    finishCooking,
    quitProgram
} from '../src/riceCooker';

function setRiceAndWaterAdded(value: boolean): void {
    // @ts-expect-error: Temporarily ignoring type-checking error to modify test-specific variable.
    riceAndWaterAdded = value;
}


describe('Rice Cooker Program', () => {
    let consoleSpy: sinon.SinonSpy;
  

    beforeEach(() => {
        consoleSpy = sinon.spy(console, 'log');
    });

    it('should have initial state "Off" and rice and water not added', () => {
        expect(riceCookerState).to.equal(STATE_OFF);
        expect(riceAndWaterAdded).to.be.false;
    });

    it('should display the initial state', () => {
        displayState();
        sinon.assert.calledWith(consoleSpy, '\nCurrent state of the rice cooker:', STATE_OFF);
    });

    it('should not allow cooking without adding rice and water', () => {
        setRiceAndWaterAdded(false);
        plugIn();
        sinon.assert.calledWith(consoleSpy, 'Error: Add rice and water before plugging in and starting cooking.');
        expect(riceAndWaterAdded).to.be.false;
    });
  
    it('should allow adding rice and water', () => {
        setRiceAndWaterAdded(false);
        setRiceAndWaterAdded(true);
        expect(riceAndWaterAdded).to.be.true;
    });

    it('should display an error when plugging in if rice and water not added', () => {
        setRiceAndWaterAdded(false);
        plugIn();
        sinon.assert.calledWith(consoleSpy, 'Error: Add rice and water before plugging in and starting cooking.');
        expect(riceCookerState).to.equal(STATE_OFF);
    });

    it('should display an error for end of cooking notification if cooking is not in progress', () => {
        setRiceAndWaterAdded(true);

        finishCooking();
        sinon.assert.calledWith(consoleSpy, 'Error: No cooking in progress.');
        expect(riceCookerState).to.equal(STATE_OFF);
    });


    it('should quit the program', () => {
        quitProgram();
        sinon.assert.calledWith(consoleSpy, '\n=======================================================================');
        sinon.assert.calledWith(consoleSpy, 'Goodbye! Thanks for using the rice cooker program.\n');
    });

    afterEach(() => {
        consoleSpy.restore();
    });
});
