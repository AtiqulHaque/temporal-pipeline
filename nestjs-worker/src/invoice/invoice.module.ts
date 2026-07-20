import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Invoice } from './invoice.entity';
import { InvoiceRepository } from './invoice.repository';
import { InvoiceService } from './invoice.service';

@Module({
  imports: [TypeOrmModule.forFeature([Invoice])],
  providers: [InvoiceRepository, InvoiceService],
  exports: [InvoiceService],
})
export class InvoiceModule {}
