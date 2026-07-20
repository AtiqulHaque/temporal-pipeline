import {
  Injectable,
  Logger,
  OnModuleDestroy,
  OnModuleInit,
} from '@nestjs/common';
import { NativeConnection, Worker } from '@temporalio/worker';
import { InvoiceService } from '../invoice/invoice.service';
import { createActivities } from './activities';

@Injectable()
export class TemporalWorkerService implements OnModuleInit, OnModuleDestroy {
  private readonly logger = new Logger(TemporalWorkerService.name);
  private worker?: Worker;

  constructor(private readonly invoiceService: InvoiceService) {}

  async onModuleInit(): Promise<void> {
    const address = process.env.TEMPORAL_ADDRESS ?? 'localhost:7233';
    const taskQueue = process.env.NESTJS_TASK_QUEUE ?? 'nestjs-task-queue';

    const connection = await NativeConnection.connect({ address });
    const activities = createActivities(this.invoiceService);

    this.worker = await Worker.create({
      connection,
      taskQueue,
      activities,
    });

    void this.worker.run().catch((error: unknown) => {
      this.logger.error('Temporal worker failed', error);
      process.exit(1);
    });

    this.logger.log(
      `NestJS worker listening on task queue "${taskQueue}" (activities: ${Object.keys(activities).join(', ')})`,
    );
  }

  async onModuleDestroy(): Promise<void> {
    if (this.worker) {
      this.worker.shutdown();
    }
  }
}
