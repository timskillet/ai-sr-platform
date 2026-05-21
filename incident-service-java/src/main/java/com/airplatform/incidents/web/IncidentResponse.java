package com.airplatform.incidents.web;

import java.time.Instant;
import java.util.UUID;

public record IncidentResponse(UUID id, String service, String severity, String status, String message, Instant createdAt) {
}
