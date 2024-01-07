import '@testing-library/jest-dom';
import { render, screen } from '@testing-library/react';
import Button from '@/app/_components/atoms/Button';

describe('Button', () => {
  it('button rendering', () => {
    const buttonElementProps = {
      onClick: () => {
        alert('button pushed!');
      },
      disabled: false,
    };

    render(
      <Button type="positive" buttonElementProps={buttonElementProps}>
        <div>Save</div>
      </Button>,
    );

    const button = screen.getByText('Save');
    expect(button).toBeInTheDocument();
  });
});
