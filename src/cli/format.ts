// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2023-Present The Pepr Authors

import { ESLint } from "eslint";
import { promises as fs } from "fs";
import { format, resolveConfig } from "prettier";

import Log from "../lib/logger";
import { RootCmd } from "./root";

export default function (program: RootCmd) {
  program
    .command("format")
    .description("Lint and format this Pepr module")
    .option("-v, --validate-only", "Do not modify files, only validate formatting")
    .action(async opts => {
      try {
        // "format:check": "eslint src && prettier src --check",
        // "format:fix": "eslint src --fix && prettier src --write"

        const eslint = new ESLint();
        const results = await eslint.lintFiles(["./**/*.ts"]);

        // Track if any files failed
        let hasFailure = false;

        results.forEach(async result => {
          const errorCount = result.fatalErrorCount + result.errorCount;
          if (errorCount > 0) {
            hasFailure = true;
          }
        });

        const formatter = await eslint.loadFormatter("stylish");
        const resultText = await formatter.format(results);

        if (resultText) {
          console.log(resultText);
        }

        // Write the fixes if not in validate-only mode
        if (!opts.validateOnly) {
          await ESLint.outputFixes(results);
        }

        // Format with Prettier
        for (const { filePath } of results) {
          const content = await fs.readFile(filePath, "utf8");
          const cfg = await resolveConfig(filePath);
          const formatted = format(content, { filepath: filePath, ...cfg });

          // If in validate-only mode, check if the file is formatted correctly
          if (opts.validateOnly) {
            if (formatted !== content) {
              hasFailure = true;
              Log.error(`File ${filePath} is not formatted correctly`);
            }
          } else {
            // Otherwise, write the formatted file
            await fs.writeFile(filePath, formatted);
          }
        }

        if (opts.validateOnly && hasFailure) {
          process.exit(1);
        }

        Log.info("Module formatted");
      } catch (e) {
        Log.debug(e);
        Log.error(e.message);
        process.exit(1);
      }
    });
}