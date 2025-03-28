/**
 * Utility for checking feature flag status
 */

/**
 * Check if a feature flag is enabled
 * @param flagName Name of the feature flag to check
 * @returns boolean indicating if the flag is enabled
 */
export function isFeatureFlagEnabled(flagName: string): boolean {
  // In a real implementation, this would check against a backend service
  // or local storage for feature flag status
  
  // For demo purposes, all flags are disabled
  return false;
}