export function formatMessageActivity(greeting: string): string {
  return `Format message: ${greeting} — processed by NestJS worker at ${new Date().toISOString()}`;
}
