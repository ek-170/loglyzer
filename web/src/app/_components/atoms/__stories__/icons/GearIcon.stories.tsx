import type { Meta, StoryObj } from '@storybook/react';

import { GearIcon } from '@/app/_components/atoms/icons/GearIcon';

const meta: Meta<typeof GearIcon> = {
  component: GearIcon,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof GearIcon>;

export const Icon: Story = {};
