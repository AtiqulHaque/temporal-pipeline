export interface CreateInvoiceDto {
  invoiceNumber: string;
  customerName: string;
  amount: number;
  currency?: string;
  status?: 'pending' | 'created' | 'failed';
  workflowId?: string;
}
