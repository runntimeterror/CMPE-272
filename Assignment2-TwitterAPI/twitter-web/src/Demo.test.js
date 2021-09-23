import Demo from './Demo'
import { render, screen } from '@testing-library/react';

test('snapshot', ()=> {
    const tree = render(<Demo />);
    expect(tree).toMatchSnapshot()
})