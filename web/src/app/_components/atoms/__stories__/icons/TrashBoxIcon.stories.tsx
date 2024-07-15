import type { Meta, StoryObj } from '@storybook/react';

import { TrashBoxIcon } from '@/app/_components/atoms/icons/TrashBoxIcon';

const meta: Meta<typeof TrashBoxIcon> = {
  component: TrashBoxIcon,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof TrashBoxIcon>;

export const Icon: Story = {};
