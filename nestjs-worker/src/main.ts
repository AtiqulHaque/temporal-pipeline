import { NestFactory } from '@nestjs/core';
import { Logger } from '@nestjs/common';
import { AppModule } from './app.module';

async function bootstrap(): Promise<void> {
  const app = await NestFactory.createApplicationContext(AppModule, {
    logger: ['log', 'error', 'warn'],
  });

  const logger = new Logger('Bootstrap');
  logger.log('NestJS Temporal worker application started');

  process.on('SIGINT', async () => {
    await app.close();
    process.exit(0);
  });

  process.on('SIGTERM', async () => {
    await app.close();
    process.exit(0);
  });
}

bootstrap().catch((error: unknown) => {
  console.error('Failed to start NestJS worker', error);
  process.exit(1);
});
