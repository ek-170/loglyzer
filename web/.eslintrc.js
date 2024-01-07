module.exports = {
  root: true,
  parser: '@typescript-eslint/parser',
  plugins: ['@typescript-eslint', 'import', 'unused-imports'],
  extends: [
    'eslint:recommended',
    'next/core-web-vitals',
    'plugin:react/recommended',
    'plugin:react-hooks/recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:tailwindcss/recommended',
    'prettier',
    'plugin:storybook/recommended',
  ],
  rules: {
    'valid-jsdoc': 'off',

    'require-jsdoc': 0,

    '@typescript-eslint/no-non-null-assertion': 'off',

    camelcase: ['error', { properties: 'never' }],

    // 'react/jsx-uses-react': 'off',
    'react/react-in-jsx-scope': 'off',

    // TypeScriptを使用している &
    // アロー関数でコンポーネントを作ると意味のないエラーが出る
    // https://github.com/yannickcr/eslint-plugin-react/issues/2353#issuecomment-534752036
    'react/prop-types': 'off',

    // ts-ignore は無視
    '@typescript-eslint/ban-ts-comment': 'off',

    // テーブルなど厳密に型をつけるのが難しい箇所があるため
    '@typescript-eslint/no-explicit-any': 'off',
    '@typescript-eslint/no-unused-vars': 'off',
    'unused-imports/no-unused-imports': 'error',
  },
  globals: {
    window: true,
  },
  settings: {
    react: {
      version: 'detect',
    },
  },
};
