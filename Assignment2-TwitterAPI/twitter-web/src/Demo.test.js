/**
 * File created by Rajat Banerjee
 */
import Demo from './Demo'
import { render } from '@testing-library/react';
import Jest from 'jest'

test('snapshot', ()=> {
    const originalError = console.error;
    console.error = jest.fn();
    const tree = render(<Demo />);
    expect(tree).toMatchSnapshot()
    console.error = originalError;
})