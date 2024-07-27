export const root = true
export const overrides = [
    {
        files: ['**/*.js', '**/*.mjs', '**/*.cjs', '**/*.ts'],
        parserOptions: {
            parser: '@typescript-eslint/parser',
        },
        extends: ['eslint:recommended', 'plugin:@typescript-eslint/recommended'],
    },
]
export const env = {
    browser: true,
    node: true,
}
export const rules = {}
