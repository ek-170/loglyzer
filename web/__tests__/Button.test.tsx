import '@testing-library/jest-dom';
import { render, screen } from '@testing-library/react';
import { Button } from '@/app/_components/atoms/buttons/Button';

describe('Button', () => {
  it('button rendering', () => {
    const buttonElementProps = {
      onClick: () => {
        alert('button pushed!');
      },
      disabled: false,
    };

    render(
      <Button color="positive" buttonElementProps={buttonElementProps}>
        <div>Save</div>
      </Button>,
    );

    const button = screen.getByText('Save');
    expect(button).toBeInTheDocument();
  });
});
