import { render } from '@testing-library/react';
import { StaticRouter } from 'react-router';
import Home from './Home';

test('snapshot', () => {
    const tree = render(<StaticRouter><Home /></StaticRouter>);
    expect(tree).toMatchSnapshot()
})
