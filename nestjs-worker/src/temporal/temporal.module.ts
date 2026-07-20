import { Module } from '@nestjs/common';
import { InvoiceModule } from '../invoice/invoice.module';
import { TemporalWorkerService } from './temporal-worker.service';

@Module({
  imports: [InvoiceModule],
  providers: [TemporalWorkerService],
})
export class TemporalModule {}
