import type { Meta, StoryObj } from '@storybook/react';

import { SearchIcon } from '@/app/_components/atoms/icons/SearchIcon';

const meta: Meta<typeof SearchIcon> = {
  component: SearchIcon,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof SearchIcon>;

export const Icon: Story = {};
