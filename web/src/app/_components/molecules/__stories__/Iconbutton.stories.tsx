import type { Meta, StoryObj } from '@storybook/react';

import { IconButton } from '@/app/_components/molecules/IconButton';
import { SearchIcon } from '../../atoms/icons/SearchIcon';

const buttonElementProps = {
  onClick: () => {
    alert('button pushed!');
  },
  disabled: false,
};

const meta: Meta<typeof IconButton> = {
  component: IconButton,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  args: {
    buttonElementProps: buttonElementProps,
    padding: ['px-3', 'py-1.5'],
  },
};

export default meta;
type Story = StoryObj<typeof IconButton>;

export const PositiveButton: Story = {
  args: {
    color: 'primary',
    startIcon: <SearchIcon width={20} height={20} color="white" />,
    children: <p className="text-white">Search</p>,
  },
};
