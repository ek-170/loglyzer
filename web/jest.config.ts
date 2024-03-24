// see more settings below
// https://nextjs.org/docs/app/building-your-application/testing/jest
import type { Config } from 'jest';
import nextJest from 'next/jest.js';

const createJestConfig = nextJest({
  dir: './',
});

const config: Config = {
  moduleDirectories: ['node_modules', '<rootDir>/'],
  coverageProvider: 'v8',
  testEnvironment: 'jsdom',
  setupFilesAfterEnv: ['<rootDir>/src/app/_test/jest.setup.ts'],
  moduleNameMapper: {
    '^@/app/(.*)$': '<rootDir>/src/app/$1',
  },
};

export default createJestConfig(config);
