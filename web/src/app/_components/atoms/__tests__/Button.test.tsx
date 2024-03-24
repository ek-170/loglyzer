import '@testing-library/jest-dom';
import { Button } from '@/app/_components/atoms';
import { render, screen } from '@/app/_test/test-utils';

describe('Button', () => {
  it('button rendering', () => {
    render(
      <Button
        color="positive"
        onClick={() => {
          alert('button pushed!');
        }}
      >
        <div>Save</div>
      </Button>,
    );

    const button = screen.getByText('Save');
    expect(button).toBeInTheDocument();
  });
});
