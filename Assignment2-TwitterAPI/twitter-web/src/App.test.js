/**
 * File created by Soham Bhattacharjee
 */
import { render, screen } from '@testing-library/react';
import App from './App';

test('renders header', () => {
  render(<App />);
  const linkElement = screen.getByText(/CMPE-272 - Assignment 2/i);
  expect(linkElement).toBeInTheDocument();
});

test('renders students', () => {
  render(<App />);
  const linkElement = screen.getByText(/Soham Bhattacharjee/i);
  expect(linkElement).toBeInTheDocument();
});

test('renders students', () => {
  render(<App />);
  const linkElement = screen.getByText(/Rajat Banerjee/i);
  expect(linkElement).toBeInTheDocument();
});
