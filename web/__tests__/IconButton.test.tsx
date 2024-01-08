import '@testing-library/jest-dom';
import { render, screen } from '@testing-library/react';
import { IconButton } from '@/app/_components/molecules/IconButton';
import { SearchIcon } from '@/app/_components/atoms/icons/SearchIcon';

describe('IconButton', () => {
  it('button rendering', () => {
    const buttonElementProps = {
      onClick: () => {
        alert('button pushed!');
      },
      disabled: false,
    };
    const startIcon = <SearchIcon width={20} height={20} color="white" />;

    render(
      <IconButton
        color="positive"
        buttonElementProps={buttonElementProps}
        startIcon={startIcon}
      >
        <div>Search</div>
      </IconButton>,
    );
    const searchIcon = screen.getByRole('img');
    expect(searchIcon).toBeInTheDocument();
  });
});
