import { Injectable, Logger } from '@nestjs/common';
import { CreateInvoiceDto } from './dto/create-invoice.dto';
import { Invoice } from './invoice.entity';
import { InvoiceRepository } from './invoice.repository';

@Injectable()
export class InvoiceService {
  private readonly logger = new Logger(InvoiceService.name);

  constructor(private readonly invoiceRepository: InvoiceRepository) {}

  async createInvoice(data: CreateInvoiceDto): Promise<Invoice> {
    this.logger.log(
      `Creating invoice ${data.invoiceNumber} for ${data.customerName}`,
    );
    const invoice = await this.invoiceRepository.create(data);
    this.logger.log(`Invoice created with id ${invoice.id}`);
    return invoice;
  }

  async getById(id: string): Promise<Invoice | null> {
    return this.invoiceRepository.findById(id);
  }
}
