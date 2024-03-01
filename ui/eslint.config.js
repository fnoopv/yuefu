import eslintJs from "@eslint/js";
import typescriptEslint from "@typescript-eslint/eslint-plugin";
import parser from "@typescript-eslint/parser";
import eslintPluginPrettierRecommended from "eslint-plugin-prettier/recommended";
import reactHooks from "eslint-plugin-react-hooks";
import reactRefresh from "eslint-plugin-react-refresh";
import globals from "globals";
import simpleImportSort from "eslint-plugin-simple-import-sort";

/**
 * @type {import("eslint").Linter.FlatConfig[]}
 */
export default [
  eslintJs.configs.recommended,
  {
    files: ["src/**/*.ts?(x)"],
    ignores: ["dist", "eslint.config.js"],
    plugins: {
      "simple-import-sort": simpleImportSort,
      "react-refresh": reactRefresh,
      "react-hooks": reactHooks,
      "@typescript-eslint": typescriptEslint,
      eslintPluginPrettierRecommended,
    },
    languageOptions: {
      parser: parser,
      sourceType: "module",
      globals: {
        ...globals.browser,
        ...globals.es2020,
      },
    },
    rules: {
      ...reactHooks.configs.recommended.rules,
      "react-refresh/only-export-components": [
        "warn",
        { allowConstantExport: true },
      ],
      "simple-import-sort/imports": "error",
      "simple-import-sort/exports": "error",
    },
  },
  eslintPluginPrettierRecommended,
];
