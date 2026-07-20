import { InvoiceService } from '../invoice/invoice.service';
import {
  CreateInvoiceInput,
  CreateInvoiceResult,
} from './createInvoice.types';

export async function createInvoiceActivity(
  invoiceService: InvoiceService,
  input: CreateInvoiceInput,
): Promise<CreateInvoiceResult> {
  const invoiceNumber =
    input.invoiceNumber ??
    `INV-${Date.now()}-${Math.floor(Math.random() * 1000)}`;

  const invoice = await invoiceService.createInvoice({
    invoiceNumber,
    customerName: input.customerName,
    amount: input.amount,
    currency: input.currency ?? 'USD',
    status: 'created',
    workflowId: input.workflowId,
  });

  return {
    id: invoice.id,
    invoiceNumber: invoice.invoiceNumber,
    customerName: invoice.customerName,
    amount: invoice.amount,
    currency: invoice.currency,
    status: invoice.status,
    workflowId: invoice.workflowId,
  };
}
