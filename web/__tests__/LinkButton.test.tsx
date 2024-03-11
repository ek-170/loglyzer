import '@testing-library/jest-dom';
import { render, screen } from '@testing-library/react';
import { LinkButton } from '@/app/_components/atoms/buttons/LinkButton';

describe('LinkButton', () => {
  it('button rendering', () => {
    render(
      <LinkButton
        color="elastic"
        href={''}
        onClick={() => {
          alert('button pushed!');
        }}
      >
        <div>Kibana</div>
      </LinkButton>,
    );

    const button = screen.getByText('Kibana');
    expect(button).toBeInTheDocument();
  });
});
