import { render, screen } from '@testing-library/react';
import CreateTweet from './Create-Tweet';

test('renders tweet button', () => {
  render(<CreateTweet />);
  const createTweetElement = screen.getByText(/Tweet/i);
  expect(createTweetElement).toBeInTheDocument()
});

test('snapshot', ()=> {
    const tree = render(<CreateTweet />);
    expect(tree).toMatchSnapshot()
})
