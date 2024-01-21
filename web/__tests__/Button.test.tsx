import '@testing-library/jest-dom';
import { render, screen } from '@testing-library/react';
import { Button } from '@/app/_components/atoms/buttons/Button';

describe('Button', () => {
  it('button rendering', () => {
    render(
      <Button color="positive" onClick={() => { alert('button pushed!'); }}>
        <div>Save</div>
      </Button>,
    );

    const button = screen.getByText('Save');
    expect(button).toBeInTheDocument();
  });
});
