import React from 'react';
import { render } from '@testing-library/react';
import App from '../pages/App';

test('renders learn react link', () => {
  const { getByText } = render(<App />);
  const text = getByText(/Welcome/i);
  expect(text).toBeInTheDocument();
});