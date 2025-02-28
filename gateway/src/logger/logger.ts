import { promises as fs } from "fs";
import { join } from "path";

const LOG_DIR = join(import.meta.dir, "logs");
const LOG_FILE = join(LOG_DIR, "app.log");


async function ensureLogDirExists() {
  try {
    await fs.mkdir(LOG_DIR, { recursive: true });
  } catch (error) {
    console.error("Failed to create log directory", error);
  }
}

ensureLogDirExists();

function formatLog(level: string, message: string, data?: any): string {
  const timestamp = new Date().toISOString();
  return `[${level}] ${timestamp} - ${message} ${data ? JSON.stringify(data) : ""}\n`;
}

async function writeLogToFile(logMessage: string) {
  try {
    await fs.appendFile(LOG_FILE, logMessage, "utf8");
  } catch (error) {
    console.error("Failed to write log to file", error);
  }
}

export class Logger {
  static async info(message: string, data?: any) {
    const logMessage = formatLog("INFO", message, data);
    console.log(logMessage.trim());
    await writeLogToFile(logMessage);
  }

  static async warn(message: string, data?: any) {
    const logMessage = formatLog("WARN", message, data);
    console.warn(logMessage.trim());
    await writeLogToFile(logMessage);
  }

  static async error(message: string, error?: any) {
    const logMessage = formatLog("ERROR", message, error);
    console.error(logMessage.trim());
    await writeLogToFile(logMessage);
  }
}
