import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { CreateInvoiceDto } from './dto/create-invoice.dto';
import { Invoice } from './invoice.entity';

@Injectable()
export class InvoiceRepository {
  constructor(
    @InjectRepository(Invoice)
    private readonly repo: Repository<Invoice>,
  ) {}

  async create(data: CreateInvoiceDto): Promise<Invoice> {
    const invoice = this.repo.create({
      invoiceNumber: data.invoiceNumber,
      customerName: data.customerName,
      amount: data.amount.toFixed(2),
      currency: data.currency ?? 'USD',
      status: data.status ?? 'created',
      workflowId: data.workflowId ?? null,
    });
    return this.repo.save(invoice);
  }

  async findById(id: string): Promise<Invoice | null> {
    return this.repo.findOne({ where: { id } });
  }

  async findByInvoiceNumber(invoiceNumber: string): Promise<Invoice | null> {
    return this.repo.findOne({ where: { invoiceNumber } });
  }
}
